package main

// DjlServing - This module contains an universal model serving implementation
// Homepage: https://github.com/deepjavalibrary/djl-serving

import (
	"fmt"
	
	"os/exec"
)

func installDjlServing() {
	// Método 1: Descargar y extraer .tar.gz
	djlserving_tar_url := "https://publish.djl.ai/djl-serving/serving-0.29.0.tar"
	djlserving_cmd_tar := exec.Command("curl", "-L", djlserving_tar_url, "-o", "package.tar.gz")
	err := djlserving_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	djlserving_zip_url := "https://publish.djl.ai/djl-serving/serving-0.29.0.tar"
	djlserving_cmd_zip := exec.Command("curl", "-L", djlserving_zip_url, "-o", "package.zip")
	err = djlserving_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	djlserving_bin_url := "https://publish.djl.ai/djl-serving/serving-0.29.0.tar"
	djlserving_cmd_bin := exec.Command("curl", "-L", djlserving_bin_url, "-o", "binary.bin")
	err = djlserving_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	djlserving_src_url := "https://publish.djl.ai/djl-serving/serving-0.29.0.tar"
	djlserving_cmd_src := exec.Command("curl", "-L", djlserving_src_url, "-o", "source.tar.gz")
	err = djlserving_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	djlserving_cmd_direct := exec.Command("./binary")
	err = djlserving_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
