package main

// YleDl - Download Yle videos from the command-line
// Homepage: https://aajanki.github.io/yle-dl/index-en.html

import (
	"fmt"
	
	"os/exec"
)

func installYleDl() {
	// Método 1: Descargar y extraer .tar.gz
	yledl_tar_url := "https://files.pythonhosted.org/packages/1f/82/cf37a73bb0c223e80484d337aed3fa0d6b855512b1b36dde9b8eb062907a/yle_dl-20240806.tar.gz"
	yledl_cmd_tar := exec.Command("curl", "-L", yledl_tar_url, "-o", "package.tar.gz")
	err := yledl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yledl_zip_url := "https://files.pythonhosted.org/packages/1f/82/cf37a73bb0c223e80484d337aed3fa0d6b855512b1b36dde9b8eb062907a/yle_dl-20240806.zip"
	yledl_cmd_zip := exec.Command("curl", "-L", yledl_zip_url, "-o", "package.zip")
	err = yledl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yledl_bin_url := "https://files.pythonhosted.org/packages/1f/82/cf37a73bb0c223e80484d337aed3fa0d6b855512b1b36dde9b8eb062907a/yle_dl-20240806.bin"
	yledl_cmd_bin := exec.Command("curl", "-L", yledl_bin_url, "-o", "binary.bin")
	err = yledl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yledl_src_url := "https://files.pythonhosted.org/packages/1f/82/cf37a73bb0c223e80484d337aed3fa0d6b855512b1b36dde9b8eb062907a/yle_dl-20240806.src.tar.gz"
	yledl_cmd_src := exec.Command("curl", "-L", yledl_src_url, "-o", "source.tar.gz")
	err = yledl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yledl_cmd_direct := exec.Command("./binary")
	err = yledl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: rtmpdump")
	exec.Command("latte", "install", "rtmpdump").Run()
}
