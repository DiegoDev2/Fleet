package main

// Libtensorflow - C interface for Google's OS library for Machine Intelligence
// Homepage: https://www.tensorflow.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibtensorflow() {
	// Método 1: Descargar y extraer .tar.gz
	libtensorflow_tar_url := "https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.17.0.tar.gz"
	libtensorflow_cmd_tar := exec.Command("curl", "-L", libtensorflow_tar_url, "-o", "package.tar.gz")
	err := libtensorflow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtensorflow_zip_url := "https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.17.0.zip"
	libtensorflow_cmd_zip := exec.Command("curl", "-L", libtensorflow_zip_url, "-o", "package.zip")
	err = libtensorflow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtensorflow_bin_url := "https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.17.0.bin"
	libtensorflow_cmd_bin := exec.Command("curl", "-L", libtensorflow_bin_url, "-o", "binary.bin")
	err = libtensorflow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtensorflow_src_url := "https://github.com/tensorflow/tensorflow/archive/refs/tags/v2.17.0.src.tar.gz"
	libtensorflow_cmd_src := exec.Command("curl", "-L", libtensorflow_src_url, "-o", "source.tar.gz")
	err = libtensorflow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtensorflow_cmd_direct := exec.Command("./binary")
	err = libtensorflow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bazelisk")
exec.Command("latte", "install", "bazelisk")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gnu-getopt")
exec.Command("latte", "install", "gnu-getopt")
}
