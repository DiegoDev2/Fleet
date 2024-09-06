package main

// OpenImageDenoise - High-performance denoising library for ray tracing
// Homepage: https://openimagedenoise.github.io

import (
	"fmt"
	
	"os/exec"
)

func installOpenImageDenoise() {
	// Método 1: Descargar y extraer .tar.gz
	openimagedenoise_tar_url := "https://github.com/OpenImageDenoise/oidn/releases/download/v2.3.0/oidn-2.3.0.src.tar.gz"
	openimagedenoise_cmd_tar := exec.Command("curl", "-L", openimagedenoise_tar_url, "-o", "package.tar.gz")
	err := openimagedenoise_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openimagedenoise_zip_url := "https://github.com/OpenImageDenoise/oidn/releases/download/v2.3.0/oidn-2.3.0.src.zip"
	openimagedenoise_cmd_zip := exec.Command("curl", "-L", openimagedenoise_zip_url, "-o", "package.zip")
	err = openimagedenoise_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openimagedenoise_bin_url := "https://github.com/OpenImageDenoise/oidn/releases/download/v2.3.0/oidn-2.3.0.src.bin"
	openimagedenoise_cmd_bin := exec.Command("curl", "-L", openimagedenoise_bin_url, "-o", "binary.bin")
	err = openimagedenoise_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openimagedenoise_src_url := "https://github.com/OpenImageDenoise/oidn/releases/download/v2.3.0/oidn-2.3.0.src.src.tar.gz"
	openimagedenoise_cmd_src := exec.Command("curl", "-L", openimagedenoise_src_url, "-o", "source.tar.gz")
	err = openimagedenoise_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openimagedenoise_cmd_direct := exec.Command("./binary")
	err = openimagedenoise_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ispc")
exec.Command("latte", "install", "ispc")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
}
