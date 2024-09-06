package main

// Oggz - Command-line tool for manipulating Ogg files
// Homepage: https://www.xiph.org/oggz/

import (
	"fmt"
	
	"os/exec"
)

func installOggz() {
	// Método 1: Descargar y extraer .tar.gz
	oggz_tar_url := "https://downloads.xiph.org/releases/liboggz/liboggz-1.1.1.tar.gz"
	oggz_cmd_tar := exec.Command("curl", "-L", oggz_tar_url, "-o", "package.tar.gz")
	err := oggz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oggz_zip_url := "https://downloads.xiph.org/releases/liboggz/liboggz-1.1.1.zip"
	oggz_cmd_zip := exec.Command("curl", "-L", oggz_zip_url, "-o", "package.zip")
	err = oggz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oggz_bin_url := "https://downloads.xiph.org/releases/liboggz/liboggz-1.1.1.bin"
	oggz_cmd_bin := exec.Command("curl", "-L", oggz_bin_url, "-o", "binary.bin")
	err = oggz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oggz_src_url := "https://downloads.xiph.org/releases/liboggz/liboggz-1.1.1.src.tar.gz"
	oggz_cmd_src := exec.Command("curl", "-L", oggz_src_url, "-o", "source.tar.gz")
	err = oggz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oggz_cmd_direct := exec.Command("./binary")
	err = oggz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
}
