package main

// SqlxCli - Command-line utility for SQLx, the Rust SQL toolkit
// Homepage: https://github.com/launchbadge/sqlx

import (
	"fmt"
	
	"os/exec"
)

func installSqlxCli() {
	// Método 1: Descargar y extraer .tar.gz
	sqlxcli_tar_url := "https://github.com/launchbadge/sqlx/archive/refs/tags/v0.8.2.tar.gz"
	sqlxcli_cmd_tar := exec.Command("curl", "-L", sqlxcli_tar_url, "-o", "package.tar.gz")
	err := sqlxcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlxcli_zip_url := "https://github.com/launchbadge/sqlx/archive/refs/tags/v0.8.2.zip"
	sqlxcli_cmd_zip := exec.Command("curl", "-L", sqlxcli_zip_url, "-o", "package.zip")
	err = sqlxcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlxcli_bin_url := "https://github.com/launchbadge/sqlx/archive/refs/tags/v0.8.2.bin"
	sqlxcli_cmd_bin := exec.Command("curl", "-L", sqlxcli_bin_url, "-o", "binary.bin")
	err = sqlxcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlxcli_src_url := "https://github.com/launchbadge/sqlx/archive/refs/tags/v0.8.2.src.tar.gz"
	sqlxcli_cmd_src := exec.Command("curl", "-L", sqlxcli_src_url, "-o", "source.tar.gz")
	err = sqlxcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlxcli_cmd_direct := exec.Command("./binary")
	err = sqlxcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
