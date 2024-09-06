package main

// Arpoison - UNIX arp cache update utility
// Homepage: http://www.arpoison.net/

import (
	"fmt"
	
	"os/exec"
)

func installArpoison() {
	// Método 1: Descargar y extraer .tar.gz
	arpoison_tar_url := "http://www.arpoison.net/arpoison-0.7.tar.gz"
	arpoison_cmd_tar := exec.Command("curl", "-L", arpoison_tar_url, "-o", "package.tar.gz")
	err := arpoison_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arpoison_zip_url := "http://www.arpoison.net/arpoison-0.7.zip"
	arpoison_cmd_zip := exec.Command("curl", "-L", arpoison_zip_url, "-o", "package.zip")
	err = arpoison_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arpoison_bin_url := "http://www.arpoison.net/arpoison-0.7.bin"
	arpoison_cmd_bin := exec.Command("curl", "-L", arpoison_bin_url, "-o", "binary.bin")
	err = arpoison_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arpoison_src_url := "http://www.arpoison.net/arpoison-0.7.src.tar.gz"
	arpoison_cmd_src := exec.Command("curl", "-L", arpoison_src_url, "-o", "source.tar.gz")
	err = arpoison_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arpoison_cmd_direct := exec.Command("./binary")
	err = arpoison_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libnet")
exec.Command("latte", "install", "libnet")
}
