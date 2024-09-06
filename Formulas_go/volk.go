package main

// Volk - Vector Optimized Library of Kernels
// Homepage: https://www.libvolk.org/

import (
	"fmt"
	
	"os/exec"
)

func installVolk() {
	// Método 1: Descargar y extraer .tar.gz
	volk_tar_url := "https://github.com/gnuradio/volk/releases/download/v3.1.2/volk-3.1.2.tar.gz"
	volk_cmd_tar := exec.Command("curl", "-L", volk_tar_url, "-o", "package.tar.gz")
	err := volk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	volk_zip_url := "https://github.com/gnuradio/volk/releases/download/v3.1.2/volk-3.1.2.zip"
	volk_cmd_zip := exec.Command("curl", "-L", volk_zip_url, "-o", "package.zip")
	err = volk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	volk_bin_url := "https://github.com/gnuradio/volk/releases/download/v3.1.2/volk-3.1.2.bin"
	volk_cmd_bin := exec.Command("curl", "-L", volk_bin_url, "-o", "binary.bin")
	err = volk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	volk_src_url := "https://github.com/gnuradio/volk/releases/download/v3.1.2/volk-3.1.2.src.tar.gz"
	volk_cmd_src := exec.Command("curl", "-L", volk_src_url, "-o", "source.tar.gz")
	err = volk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	volk_cmd_direct := exec.Command("./binary")
	err = volk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cpu_features")
exec.Command("latte", "install", "cpu_features")
	fmt.Println("Instalando dependencia: orc")
exec.Command("latte", "install", "orc")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
