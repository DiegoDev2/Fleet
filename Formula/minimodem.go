package main

// Minimodem - General-purpose software audio FSK modem
// Homepage: http://www.whence.com/minimodem/

import (
	"fmt"
	
	"os/exec"
)

func installMinimodem() {
	// Método 1: Descargar y extraer .tar.gz
	minimodem_tar_url := "http://www.whence.com/minimodem/minimodem-0.24.tar.gz"
	minimodem_cmd_tar := exec.Command("curl", "-L", minimodem_tar_url, "-o", "package.tar.gz")
	err := minimodem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minimodem_zip_url := "http://www.whence.com/minimodem/minimodem-0.24.zip"
	minimodem_cmd_zip := exec.Command("curl", "-L", minimodem_zip_url, "-o", "package.zip")
	err = minimodem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minimodem_bin_url := "http://www.whence.com/minimodem/minimodem-0.24.bin"
	minimodem_cmd_bin := exec.Command("curl", "-L", minimodem_bin_url, "-o", "binary.bin")
	err = minimodem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minimodem_src_url := "http://www.whence.com/minimodem/minimodem-0.24.src.tar.gz"
	minimodem_cmd_src := exec.Command("curl", "-L", minimodem_src_url, "-o", "source.tar.gz")
	err = minimodem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minimodem_cmd_direct := exec.Command("./binary")
	err = minimodem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
}
