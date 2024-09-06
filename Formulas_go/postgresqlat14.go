package main

// PostgresqlAT14 - Object-relational database system
// Homepage: https://www.postgresql.org/

import (
	"fmt"
	
	"os/exec"
)

func installPostgresqlAT14() {
	// Método 1: Descargar y extraer .tar.gz
	postgresqlat14_tar_url := "https://ftp.postgresql.org/pub/source/v14.13/postgresql-14.13.tar.bz2"
	postgresqlat14_cmd_tar := exec.Command("curl", "-L", postgresqlat14_tar_url, "-o", "package.tar.gz")
	err := postgresqlat14_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	postgresqlat14_zip_url := "https://ftp.postgresql.org/pub/source/v14.13/postgresql-14.13.tar.bz2"
	postgresqlat14_cmd_zip := exec.Command("curl", "-L", postgresqlat14_zip_url, "-o", "package.zip")
	err = postgresqlat14_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	postgresqlat14_bin_url := "https://ftp.postgresql.org/pub/source/v14.13/postgresql-14.13.tar.bz2"
	postgresqlat14_cmd_bin := exec.Command("curl", "-L", postgresqlat14_bin_url, "-o", "binary.bin")
	err = postgresqlat14_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	postgresqlat14_src_url := "https://ftp.postgresql.org/pub/source/v14.13/postgresql-14.13.tar.bz2"
	postgresqlat14_cmd_src := exec.Command("curl", "-L", postgresqlat14_src_url, "-o", "source.tar.gz")
	err = postgresqlat14_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	postgresqlat14_cmd_direct := exec.Command("./binary")
	err = postgresqlat14_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: krb5")
exec.Command("latte", "install", "krb5")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: linux-pam")
exec.Command("latte", "install", "linux-pam")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
