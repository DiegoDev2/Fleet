package main

// Vectorscan - High-performance regular expression matching library
// Homepage: https://github.com/VectorCamp/vectorscan

import (
	"fmt"
	
	"os/exec"
)

func installVectorscan() {
	// Método 1: Descargar y extraer .tar.gz
	vectorscan_tar_url := "https://github.com/VectorCamp/vectorscan/archive/refs/tags/vectorscan/5.4.11.tar.gz"
	vectorscan_cmd_tar := exec.Command("curl", "-L", vectorscan_tar_url, "-o", "package.tar.gz")
	err := vectorscan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vectorscan_zip_url := "https://github.com/VectorCamp/vectorscan/archive/refs/tags/vectorscan/5.4.11.zip"
	vectorscan_cmd_zip := exec.Command("curl", "-L", vectorscan_zip_url, "-o", "package.zip")
	err = vectorscan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vectorscan_bin_url := "https://github.com/VectorCamp/vectorscan/archive/refs/tags/vectorscan/5.4.11.bin"
	vectorscan_cmd_bin := exec.Command("curl", "-L", vectorscan_bin_url, "-o", "binary.bin")
	err = vectorscan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vectorscan_src_url := "https://github.com/VectorCamp/vectorscan/archive/refs/tags/vectorscan/5.4.11.src.tar.gz"
	vectorscan_cmd_src := exec.Command("curl", "-L", vectorscan_src_url, "-o", "source.tar.gz")
	err = vectorscan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vectorscan_cmd_direct := exec.Command("./binary")
	err = vectorscan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ragel")
exec.Command("latte", "install", "ragel")
}
