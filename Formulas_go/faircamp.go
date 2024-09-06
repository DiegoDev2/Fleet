package main

// Faircamp - Static site generator for audio producers
// Homepage: https://codeberg.org/simonrepp/faircamp

import (
	"fmt"
	
	"os/exec"
)

func installFaircamp() {
	// Método 1: Descargar y extraer .tar.gz
	faircamp_tar_url := "https://codeberg.org/simonrepp/faircamp/archive/0.15.1.tar.gz"
	faircamp_cmd_tar := exec.Command("curl", "-L", faircamp_tar_url, "-o", "package.tar.gz")
	err := faircamp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	faircamp_zip_url := "https://codeberg.org/simonrepp/faircamp/archive/0.15.1.zip"
	faircamp_cmd_zip := exec.Command("curl", "-L", faircamp_zip_url, "-o", "package.zip")
	err = faircamp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	faircamp_bin_url := "https://codeberg.org/simonrepp/faircamp/archive/0.15.1.bin"
	faircamp_cmd_bin := exec.Command("curl", "-L", faircamp_bin_url, "-o", "binary.bin")
	err = faircamp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	faircamp_src_url := "https://codeberg.org/simonrepp/faircamp/archive/0.15.1.src.tar.gz"
	faircamp_cmd_src := exec.Command("curl", "-L", faircamp_src_url, "-o", "source.tar.gz")
	err = faircamp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	faircamp_cmd_direct := exec.Command("./binary")
	err = faircamp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: opus")
exec.Command("latte", "install", "opus")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: vips")
exec.Command("latte", "install", "vips")
}
