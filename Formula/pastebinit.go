package main

// Pastebinit - Send things to pastebin from the command-line
// Homepage: https://launchpad.net/pastebinit

import (
	"fmt"
	
	"os/exec"
)

func installPastebinit() {
	// Método 1: Descargar y extraer .tar.gz
	pastebinit_tar_url := "https://launchpad.net/pastebinit/trunk/1.5/+download/pastebinit-1.5.tar.gz"
	pastebinit_cmd_tar := exec.Command("curl", "-L", pastebinit_tar_url, "-o", "package.tar.gz")
	err := pastebinit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pastebinit_zip_url := "https://launchpad.net/pastebinit/trunk/1.5/+download/pastebinit-1.5.zip"
	pastebinit_cmd_zip := exec.Command("curl", "-L", pastebinit_zip_url, "-o", "package.zip")
	err = pastebinit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pastebinit_bin_url := "https://launchpad.net/pastebinit/trunk/1.5/+download/pastebinit-1.5.bin"
	pastebinit_cmd_bin := exec.Command("curl", "-L", pastebinit_bin_url, "-o", "binary.bin")
	err = pastebinit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pastebinit_src_url := "https://launchpad.net/pastebinit/trunk/1.5/+download/pastebinit-1.5.src.tar.gz"
	pastebinit_cmd_src := exec.Command("curl", "-L", pastebinit_src_url, "-o", "source.tar.gz")
	err = pastebinit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pastebinit_cmd_direct := exec.Command("./binary")
	err = pastebinit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook2x")
	exec.Command("latte", "install", "docbook2x").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
