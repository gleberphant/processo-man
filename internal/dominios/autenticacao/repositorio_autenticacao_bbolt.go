package autenticacao

// tokens e permissões estão armazenadas no Bolt Database
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gleberphant/ProcessoMan/internal/entidades"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

type RepositorioTokenBolt struct {
	Conn *bolt.DB
}

// NovoRepositorioTokenBolt cria uma nova instância do repositório de tokens e estabelece a conexão.
func NovoRepositorioTokenBolt(conn *bolt.DB) *RepositorioTokenBolt {

	repo := RepositorioTokenBolt{
		Conn: conn,
	}

	return &repo

}

// Fechar encerra a conexão ativa com o banco de dados.
func (r *RepositorioTokenBolt) Fechar() {
	r.Conn.Close()
}

func (r *RepositorioTokenBolt) VerificarPermissaoPerfil(chaveRota string, perfil string) bool {
	db := r.Conn
	err := db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte("permissoes"))
		if bucket == nil {
			return fmt.Errorf("bucket não encontrada %s", chaveRota)
		}
		bytesPerfis := bucket.Get([]byte(chaveRota))

		if bytesPerfis == nil {
			return fmt.Errorf("rota não encontrada %s", chaveRota)
		}

		var perfis map[string]bool

		err := json.Unmarshal(bytesPerfis, &perfis)

		if err != nil {
			return err
		}

		if !perfis[perfil] {
			return fmt.Errorf("perfil não autorizado %v", perfis[perfil])

		}
		return nil

	})

	if err != nil {
		log.Printf("%v", err)
		return false
	}

	return true

}

// Criar insere um novo token na tabela de tokens.
func (r *RepositorioTokenBolt) Criar(token *entidades.Token) (*entidades.Token, error) {
	db := r.Conn

	token.DataCriacao = time.Now()

	err := db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists([]byte("tokens"))

		if err != nil {
			return err
		}

		bytesToken, _ := json.MarshalIndent(token, "", "	")

		bucket.Put([]byte(token.UUID.String()), bytesToken)

		return nil

	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// Deletar remove todos os tokens associados a um UUID de usuário específico.
func (r *RepositorioTokenBolt) DeletarPorUsuarioUUID(UsuarioUUID uuid.UUID) error {

	db := r.Conn

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("tokens"))

		if bucket == nil {
			return nil //errors.New("bucket nao encontrado")
		}

		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {

			var token entidades.Token

			if err := json.Unmarshal(v, &token); err != nil {
				return err
			}

			if token.UsuarioUUID == UsuarioUUID {

				err := c.Delete()

				if err != nil {
					return err
				}
			}

		}

		return nil

	})

	if err != nil {
		return err
	}

	return nil
}

// Ver busca os detalhes de um token específico através do seu UUID.
func (r *RepositorioTokenBolt) BuscarPorUUID(UUID uuid.UUID) (*entidades.Token, error) {

	db := r.Conn
	token := entidades.Token{}
	err := db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte("tokens"))

		if bucket == nil {
			return errors.New("bucket de token inexistente")
		}

		bytesToken := bucket.Get([]byte(UUID.String()))

		if bytesToken == nil {
			return errors.New("token inexistente")
		}

		json.Unmarshal(bytesToken, &token)

		return nil

	})

	if err != nil {
		return nil, err
	}

	return &token, nil

}
