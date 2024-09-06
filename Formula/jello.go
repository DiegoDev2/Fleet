package main

// Jello - Filter JSON and JSON Lines data with Python syntax
// Homepage: https://github.com/kellyjonbrazil/jello

import (
	"fmt"
	
	"os/exec"
)

func installJello() {
	// Método 1: Descargar y extraer .tar.gz
	jello_tar_url := "https://files.pythonhosted.org/packages/8a/1d/25e13e337f0c5c8076a4fc42db02b726529b611a69d816b71f8d591cf0f5/jello-1.6.0.tar.gz"
	jello_cmd_tar := exec.Command("curl", "-L", jello_tar_url, "-o", "package.tar.gz")
	err := jello_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jello_zip_url := "https://files.pythonhosted.org/packages/8a/1d/25e13e337f0c5c8076a4fc42db02b726529b611a69d816b71f8d591cf0f5/jello-1.6.0.zip"
	jello_cmd_zip := exec.Command("curl", "-L", jello_zip_url, "-o", "package.zip")
	err = jello_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jello_bin_url := "https://files.pythonhosted.org/packages/8a/1d/25e13e337f0c5c8076a4fc42db02b726529b611a69d816b71f8d591cf0f5/jello-1.6.0.bin"
	jello_cmd_bin := exec.Command("curl", "-L", jello_bin_url, "-o", "binary.bin")
	err = jello_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jello_src_url := "https://files.pythonhosted.org/packages/8a/1d/25e13e337f0c5c8076a4fc42db02b726529b611a69d816b71f8d591cf0f5/jello-1.6.0.src.tar.gz"
	jello_cmd_src := exec.Command("curl", "-L", jello_src_url, "-o", "source.tar.gz")
	err = jello_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jello_cmd_direct := exec.Command("./binary")
	err = jello_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
