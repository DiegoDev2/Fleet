package main

// Pgbouncer - Lightweight connection pooler for PostgreSQL
// Homepage: https://www.pgbouncer.org/

import (
	"fmt"
	
	"os/exec"
)

func installPgbouncer() {
	// Método 1: Descargar y extraer .tar.gz
	pgbouncer_tar_url := "https://www.pgbouncer.org/downloads/files/1.23.1/pgbouncer-1.23.1.tar.gz"
	pgbouncer_cmd_tar := exec.Command("curl", "-L", pgbouncer_tar_url, "-o", "package.tar.gz")
	err := pgbouncer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgbouncer_zip_url := "https://www.pgbouncer.org/downloads/files/1.23.1/pgbouncer-1.23.1.zip"
	pgbouncer_cmd_zip := exec.Command("curl", "-L", pgbouncer_zip_url, "-o", "package.zip")
	err = pgbouncer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgbouncer_bin_url := "https://www.pgbouncer.org/downloads/files/1.23.1/pgbouncer-1.23.1.bin"
	pgbouncer_cmd_bin := exec.Command("curl", "-L", pgbouncer_bin_url, "-o", "binary.bin")
	err = pgbouncer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgbouncer_src_url := "https://www.pgbouncer.org/downloads/files/1.23.1/pgbouncer-1.23.1.src.tar.gz"
	pgbouncer_cmd_src := exec.Command("curl", "-L", pgbouncer_src_url, "-o", "source.tar.gz")
	err = pgbouncer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgbouncer_cmd_direct := exec.Command("./binary")
	err = pgbouncer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
