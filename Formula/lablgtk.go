package main

// Lablgtk - Objective Caml interface to gtk+
// Homepage: https://github.com/garrigue/lablgtk

import (
	"fmt"
	
	"os/exec"
)

func installLablgtk() {
	// Método 1: Descargar y extraer .tar.gz
	lablgtk_tar_url := "https://github.com/garrigue/lablgtk/archive/refs/tags/2.18.12.tar.gz"
	lablgtk_cmd_tar := exec.Command("curl", "-L", lablgtk_tar_url, "-o", "package.tar.gz")
	err := lablgtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lablgtk_zip_url := "https://github.com/garrigue/lablgtk/archive/refs/tags/2.18.12.zip"
	lablgtk_cmd_zip := exec.Command("curl", "-L", lablgtk_zip_url, "-o", "package.zip")
	err = lablgtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lablgtk_bin_url := "https://github.com/garrigue/lablgtk/archive/refs/tags/2.18.12.bin"
	lablgtk_cmd_bin := exec.Command("curl", "-L", lablgtk_bin_url, "-o", "binary.bin")
	err = lablgtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lablgtk_src_url := "https://github.com/garrigue/lablgtk/archive/refs/tags/2.18.12.src.tar.gz"
	lablgtk_cmd_src := exec.Command("curl", "-L", lablgtk_src_url, "-o", "source.tar.gz")
	err = lablgtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lablgtk_cmd_direct := exec.Command("./binary")
	err = lablgtk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
	fmt.Println("Instalando dependencia: gtksourceview")
	exec.Command("latte", "install", "gtksourceview").Run()
	fmt.Println("Instalando dependencia: librsvg")
	exec.Command("latte", "install", "librsvg").Run()
	fmt.Println("Instalando dependencia: ocaml")
	exec.Command("latte", "install", "ocaml").Run()
}
