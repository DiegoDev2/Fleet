package main

// Gnuradio - SDK for signal processing blocks to implement software radios
// Homepage: https://gnuradio.org/

import (
	"fmt"
	
	"os/exec"
)

func installGnuradio() {
	// Método 1: Descargar y extraer .tar.gz
	gnuradio_tar_url := "https://github.com/gnuradio/gnuradio/archive/refs/tags/v3.10.9.2.tar.gz"
	gnuradio_cmd_tar := exec.Command("curl", "-L", gnuradio_tar_url, "-o", "package.tar.gz")
	err := gnuradio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnuradio_zip_url := "https://github.com/gnuradio/gnuradio/archive/refs/tags/v3.10.9.2.zip"
	gnuradio_cmd_zip := exec.Command("curl", "-L", gnuradio_zip_url, "-o", "package.zip")
	err = gnuradio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnuradio_bin_url := "https://github.com/gnuradio/gnuradio/archive/refs/tags/v3.10.9.2.bin"
	gnuradio_cmd_bin := exec.Command("curl", "-L", gnuradio_bin_url, "-o", "binary.bin")
	err = gnuradio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnuradio_src_url := "https://github.com/gnuradio/gnuradio/archive/refs/tags/v3.10.9.2.src.tar.gz"
	gnuradio_cmd_src := exec.Command("curl", "-L", gnuradio_src_url, "-o", "source.tar.gz")
	err = gnuradio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnuradio_cmd_direct := exec.Command("./binary")
	err = gnuradio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: pybind11")
exec.Command("latte", "install", "pybind11")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cppzmq")
exec.Command("latte", "install", "cppzmq")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: log4cpp")
exec.Command("latte", "install", "log4cpp")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
	fmt.Println("Instalando dependencia: pygobject3")
exec.Command("latte", "install", "pygobject3")
	fmt.Println("Instalando dependencia: pyqt@5")
exec.Command("latte", "install", "pyqt@5")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: qwt-qt5")
exec.Command("latte", "install", "qwt-qt5")
	fmt.Println("Instalando dependencia: soapyrtlsdr")
exec.Command("latte", "install", "soapyrtlsdr")
	fmt.Println("Instalando dependencia: soapysdr")
exec.Command("latte", "install", "soapysdr")
	fmt.Println("Instalando dependencia: spdlog")
exec.Command("latte", "install", "spdlog")
	fmt.Println("Instalando dependencia: uhd")
exec.Command("latte", "install", "uhd")
	fmt.Println("Instalando dependencia: volk")
exec.Command("latte", "install", "volk")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
