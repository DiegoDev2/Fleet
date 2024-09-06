package main

// Bozohttpd - Small and secure http version 1.1 server
// Homepage: http://www.eterna.com.au/bozohttpd/

import (
	"fmt"
	
	"os/exec"
)

func installBozohttpd() {
	// Método 1: Descargar y extraer .tar.gz
	bozohttpd_tar_url := "http://www.eterna.com.au/bozohttpd/bozohttpd-20220517.tar.bz2"
	bozohttpd_cmd_tar := exec.Command("curl", "-L", bozohttpd_tar_url, "-o", "package.tar.gz")
	err := bozohttpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bozohttpd_zip_url := "http://www.eterna.com.au/bozohttpd/bozohttpd-20220517.tar.bz2"
	bozohttpd_cmd_zip := exec.Command("curl", "-L", bozohttpd_zip_url, "-o", "package.zip")
	err = bozohttpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bozohttpd_bin_url := "http://www.eterna.com.au/bozohttpd/bozohttpd-20220517.tar.bz2"
	bozohttpd_cmd_bin := exec.Command("curl", "-L", bozohttpd_bin_url, "-o", "binary.bin")
	err = bozohttpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bozohttpd_src_url := "http://www.eterna.com.au/bozohttpd/bozohttpd-20220517.tar.bz2"
	bozohttpd_cmd_src := exec.Command("curl", "-L", bozohttpd_src_url, "-o", "source.tar.gz")
	err = bozohttpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bozohttpd_cmd_direct := exec.Command("./binary")
	err = bozohttpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
