package main

// Diesel - Command-line tool for Rust ORM Diesel
// Homepage: https://diesel.rs

import (
	"fmt"
	
	"os/exec"
)

func installDiesel() {
	// Método 1: Descargar y extraer .tar.gz
	diesel_tar_url := "https://github.com/diesel-rs/diesel/archive/refs/tags/v2.2.4.tar.gz"
	diesel_cmd_tar := exec.Command("curl", "-L", diesel_tar_url, "-o", "package.tar.gz")
	err := diesel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diesel_zip_url := "https://github.com/diesel-rs/diesel/archive/refs/tags/v2.2.4.zip"
	diesel_cmd_zip := exec.Command("curl", "-L", diesel_zip_url, "-o", "package.zip")
	err = diesel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diesel_bin_url := "https://github.com/diesel-rs/diesel/archive/refs/tags/v2.2.4.bin"
	diesel_cmd_bin := exec.Command("curl", "-L", diesel_bin_url, "-o", "binary.bin")
	err = diesel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diesel_src_url := "https://github.com/diesel-rs/diesel/archive/refs/tags/v2.2.4.src.tar.gz"
	diesel_cmd_src := exec.Command("curl", "-L", diesel_src_url, "-o", "source.tar.gz")
	err = diesel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diesel_cmd_direct := exec.Command("./binary")
	err = diesel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: mysql-client@8.4")
	exec.Command("latte", "install", "mysql-client@8.4").Run()
}
