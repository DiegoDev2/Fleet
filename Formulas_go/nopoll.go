package main

// Nopoll - Open-source C WebSocket toolkit
// Homepage: https://www.aspl.es/nopoll/

import (
	"fmt"
	
	"os/exec"
)

func installNopoll() {
	// Método 1: Descargar y extraer .tar.gz
	nopoll_tar_url := "https://www.aspl.es/nopoll/downloads/nopoll-0.4.8.b429.tar.gz"
	nopoll_cmd_tar := exec.Command("curl", "-L", nopoll_tar_url, "-o", "package.tar.gz")
	err := nopoll_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nopoll_zip_url := "https://www.aspl.es/nopoll/downloads/nopoll-0.4.8.b429.zip"
	nopoll_cmd_zip := exec.Command("curl", "-L", nopoll_zip_url, "-o", "package.zip")
	err = nopoll_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nopoll_bin_url := "https://www.aspl.es/nopoll/downloads/nopoll-0.4.8.b429.bin"
	nopoll_cmd_bin := exec.Command("curl", "-L", nopoll_bin_url, "-o", "binary.bin")
	err = nopoll_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nopoll_src_url := "https://www.aspl.es/nopoll/downloads/nopoll-0.4.8.b429.src.tar.gz"
	nopoll_cmd_src := exec.Command("curl", "-L", nopoll_src_url, "-o", "source.tar.gz")
	err = nopoll_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nopoll_cmd_direct := exec.Command("./binary")
	err = nopoll_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
