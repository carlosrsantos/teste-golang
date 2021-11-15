package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua General", "Rua"},
		{"Avenida Lincoln", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"RUA DOS BOBOS", "Rua"},
		{"Avenida DOS VIAJANTES", "Avenida"},
		{"Estrada de terra", "Estrada"},
		{"Estrada S/NOME", "Estrada"},
		{"Praça da Independênca", "Praça"},
		{"Praça DA FÉ", "Praça"},
		/* 	{"Travessa Marechal Deodoro", "Tipo Inválido"}, */
		/* {"", "Tipo Inválido"}, */
	}

	for _, cenario := range cenarioDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}
