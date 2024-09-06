package main

// PostgresqlAT10 - Object-relational database system
// Homepage: https://www.postgresql.org/

import (
	"fmt"
	
	"os/exec"
)

func installPostgresqlAT10() {
	// Método 1: Descargar y extraer .tar.gz
	postgresqlat10_tar_url := "https://ftp.postgresql.org/pub/source/v10.22/postgresql-10.22.tar.bz2"
	postgresqlat10_cmd_tar := exec.Command("curl", "-L", postgresqlat10_tar_url, "-o", "package.tar.gz")
	err := postgresqlat10_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	postgresqlat10_zip_url := "https://ftp.postgresql.org/pub/source/v10.22/postgresql-10.22.tar.bz2"
	postgresqlat10_cmd_zip := exec.Command("curl", "-L", postgresqlat10_zip_url, "-o", "package.zip")
	err = postgresqlat10_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	postgresqlat10_bin_url := "https://ftp.postgresql.org/pub/source/v10.22/postgresql-10.22.tar.bz2"
	postgresqlat10_cmd_bin := exec.Command("curl", "-L", postgresqlat10_bin_url, "-o", "binary.bin")
	err = postgresqlat10_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	postgresqlat10_src_url := "https://ftp.postgresql.org/pub/source/v10.22/postgresql-10.22.tar.bz2"
	postgresqlat10_cmd_src := exec.Command("curl", "-L", postgresqlat10_src_url, "-o", "source.tar.gz")
	err = postgresqlat10_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	postgresqlat10_cmd_direct := exec.Command("./binary")
	err = postgresqlat10_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
