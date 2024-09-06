package main

// Coturn - Free open source implementation of TURN and STUN Server
// Homepage: https://github.com/coturn/coturn

import (
	"fmt"
	
	"os/exec"
)

func installCoturn() {
	// Método 1: Descargar y extraer .tar.gz
	coturn_tar_url := "https://github.com/coturn/coturn/archive/refs/tags/4.6.2.tar.gz"
	coturn_cmd_tar := exec.Command("curl", "-L", coturn_tar_url, "-o", "package.tar.gz")
	err := coturn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coturn_zip_url := "https://github.com/coturn/coturn/archive/refs/tags/4.6.2.zip"
	coturn_cmd_zip := exec.Command("curl", "-L", coturn_zip_url, "-o", "package.zip")
	err = coturn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coturn_bin_url := "https://github.com/coturn/coturn/archive/refs/tags/4.6.2.bin"
	coturn_cmd_bin := exec.Command("curl", "-L", coturn_bin_url, "-o", "binary.bin")
	err = coturn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coturn_src_url := "https://github.com/coturn/coturn/archive/refs/tags/4.6.2.src.tar.gz"
	coturn_cmd_src := exec.Command("curl", "-L", coturn_src_url, "-o", "source.tar.gz")
	err = coturn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coturn_cmd_direct := exec.Command("./binary")
	err = coturn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: hiredis")
	exec.Command("latte", "install", "hiredis").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
