package main

// Gifsicle - GIF image/animation creator/editor
// Homepage: https://www.lcdf.org/gifsicle/

import (
	"fmt"
	
	"os/exec"
)

func installGifsicle() {
	// Método 1: Descargar y extraer .tar.gz
	gifsicle_tar_url := "https://www.lcdf.org/gifsicle/gifsicle-1.95.tar.gz"
	gifsicle_cmd_tar := exec.Command("curl", "-L", gifsicle_tar_url, "-o", "package.tar.gz")
	err := gifsicle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gifsicle_zip_url := "https://www.lcdf.org/gifsicle/gifsicle-1.95.zip"
	gifsicle_cmd_zip := exec.Command("curl", "-L", gifsicle_zip_url, "-o", "package.zip")
	err = gifsicle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gifsicle_bin_url := "https://www.lcdf.org/gifsicle/gifsicle-1.95.bin"
	gifsicle_cmd_bin := exec.Command("curl", "-L", gifsicle_bin_url, "-o", "binary.bin")
	err = gifsicle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gifsicle_src_url := "https://www.lcdf.org/gifsicle/gifsicle-1.95.src.tar.gz"
	gifsicle_cmd_src := exec.Command("curl", "-L", gifsicle_src_url, "-o", "source.tar.gz")
	err = gifsicle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gifsicle_cmd_direct := exec.Command("./binary")
	err = gifsicle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
