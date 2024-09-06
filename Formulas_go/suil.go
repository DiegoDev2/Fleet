package main

// Suil - Lightweight C library for loading and wrapping LV2 plugin UIs
// Homepage: https://drobilla.net/software/suil.html

import (
	"fmt"
	
	"os/exec"
)

func installSuil() {
	// Método 1: Descargar y extraer .tar.gz
	suil_tar_url := "https://download.drobilla.net/suil-0.10.20.tar.xz"
	suil_cmd_tar := exec.Command("curl", "-L", suil_tar_url, "-o", "package.tar.gz")
	err := suil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	suil_zip_url := "https://download.drobilla.net/suil-0.10.20.tar.xz"
	suil_cmd_zip := exec.Command("curl", "-L", suil_zip_url, "-o", "package.zip")
	err = suil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	suil_bin_url := "https://download.drobilla.net/suil-0.10.20.tar.xz"
	suil_cmd_bin := exec.Command("curl", "-L", suil_bin_url, "-o", "binary.bin")
	err = suil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	suil_src_url := "https://download.drobilla.net/suil-0.10.20.tar.xz"
	suil_cmd_src := exec.Command("curl", "-L", suil_src_url, "-o", "source.tar.gz")
	err = suil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	suil_cmd_direct := exec.Command("./binary")
	err = suil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: lv2")
exec.Command("latte", "install", "lv2")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
}
