package main

// VowpalWabbit - Online learning algorithm
// Homepage: https://github.com/VowpalWabbit/vowpal_wabbit

import (
	"fmt"
	
	"os/exec"
)

func installVowpalWabbit() {
	// Método 1: Descargar y extraer .tar.gz
	vowpalwabbit_tar_url := "https://github.com/VowpalWabbit/vowpal_wabbit/archive/refs/tags/9.10.0.tar.gz"
	vowpalwabbit_cmd_tar := exec.Command("curl", "-L", vowpalwabbit_tar_url, "-o", "package.tar.gz")
	err := vowpalwabbit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vowpalwabbit_zip_url := "https://github.com/VowpalWabbit/vowpal_wabbit/archive/refs/tags/9.10.0.zip"
	vowpalwabbit_cmd_zip := exec.Command("curl", "-L", vowpalwabbit_zip_url, "-o", "package.zip")
	err = vowpalwabbit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vowpalwabbit_bin_url := "https://github.com/VowpalWabbit/vowpal_wabbit/archive/refs/tags/9.10.0.bin"
	vowpalwabbit_cmd_bin := exec.Command("curl", "-L", vowpalwabbit_bin_url, "-o", "binary.bin")
	err = vowpalwabbit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vowpalwabbit_src_url := "https://github.com/VowpalWabbit/vowpal_wabbit/archive/refs/tags/9.10.0.src.tar.gz"
	vowpalwabbit_cmd_src := exec.Command("curl", "-L", vowpalwabbit_src_url, "-o", "source.tar.gz")
	err = vowpalwabbit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vowpalwabbit_cmd_direct := exec.Command("./binary")
	err = vowpalwabbit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rapidjson")
	exec.Command("latte", "install", "rapidjson").Run()
	fmt.Println("Instalando dependencia: spdlog")
	exec.Command("latte", "install", "spdlog").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: sse2neon")
	exec.Command("latte", "install", "sse2neon").Run()
}
