package main

// Ncdu - NCurses Disk Usage
// Homepage: https://dev.yorhel.nl/ncdu

import (
	"fmt"
	
	"os/exec"
)

func installNcdu() {
	// Método 1: Descargar y extraer .tar.gz
	ncdu_tar_url := "https://dev.yorhel.nl/download/ncdu-2.5.tar.gz"
	ncdu_cmd_tar := exec.Command("curl", "-L", ncdu_tar_url, "-o", "package.tar.gz")
	err := ncdu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncdu_zip_url := "https://dev.yorhel.nl/download/ncdu-2.5.zip"
	ncdu_cmd_zip := exec.Command("curl", "-L", ncdu_zip_url, "-o", "package.zip")
	err = ncdu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncdu_bin_url := "https://dev.yorhel.nl/download/ncdu-2.5.bin"
	ncdu_cmd_bin := exec.Command("curl", "-L", ncdu_bin_url, "-o", "binary.bin")
	err = ncdu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncdu_src_url := "https://dev.yorhel.nl/download/ncdu-2.5.src.tar.gz"
	ncdu_cmd_src := exec.Command("curl", "-L", ncdu_src_url, "-o", "source.tar.gz")
	err = ncdu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncdu_cmd_direct := exec.Command("./binary")
	err = ncdu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: zig")
exec.Command("latte", "install", "zig")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
}
