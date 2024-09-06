package main

// Postgrest - Serves a fully RESTful API from any existing PostgreSQL database
// Homepage: https://github.com/PostgREST/postgrest

import (
	"fmt"
	
	"os/exec"
)

func installPostgrest() {
	// Método 1: Descargar y extraer .tar.gz
	postgrest_tar_url := "https://github.com/PostgREST/postgrest/archive/refs/tags/v12.2.3.tar.gz"
	postgrest_cmd_tar := exec.Command("curl", "-L", postgrest_tar_url, "-o", "package.tar.gz")
	err := postgrest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	postgrest_zip_url := "https://github.com/PostgREST/postgrest/archive/refs/tags/v12.2.3.zip"
	postgrest_cmd_zip := exec.Command("curl", "-L", postgrest_zip_url, "-o", "package.zip")
	err = postgrest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	postgrest_bin_url := "https://github.com/PostgREST/postgrest/archive/refs/tags/v12.2.3.bin"
	postgrest_cmd_bin := exec.Command("curl", "-L", postgrest_bin_url, "-o", "binary.bin")
	err = postgrest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	postgrest_src_url := "https://github.com/PostgREST/postgrest/archive/refs/tags/v12.2.3.src.tar.gz"
	postgrest_cmd_src := exec.Command("curl", "-L", postgrest_src_url, "-o", "source.tar.gz")
	err = postgrest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	postgrest_cmd_direct := exec.Command("./binary")
	err = postgrest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
exec.Command("latte", "install", "cabal-install")
	fmt.Println("Instalando dependencia: ghc@9.8")
exec.Command("latte", "install", "ghc@9.8")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
}
