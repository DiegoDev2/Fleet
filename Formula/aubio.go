package main

// Aubio - Extract annotations from audio signals
// Homepage: https://github.com/aubio/aubio

import (
	"fmt"
	
	"os/exec"
)

func installAubio() {
	// Método 1: Descargar y extraer .tar.gz
	aubio_tar_url := "https://sources.buildroot.net/aubio/aubio-0.4.9.tar.bz2"
	aubio_cmd_tar := exec.Command("curl", "-L", aubio_tar_url, "-o", "package.tar.gz")
	err := aubio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aubio_zip_url := "https://sources.buildroot.net/aubio/aubio-0.4.9.tar.bz2"
	aubio_cmd_zip := exec.Command("curl", "-L", aubio_zip_url, "-o", "package.zip")
	err = aubio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aubio_bin_url := "https://sources.buildroot.net/aubio/aubio-0.4.9.tar.bz2"
	aubio_cmd_bin := exec.Command("curl", "-L", aubio_bin_url, "-o", "binary.bin")
	err = aubio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aubio_src_url := "https://sources.buildroot.net/aubio/aubio-0.4.9.tar.bz2"
	aubio_cmd_src := exec.Command("curl", "-L", aubio_src_url, "-o", "source.tar.gz")
	err = aubio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aubio_cmd_direct := exec.Command("./binary")
	err = aubio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
