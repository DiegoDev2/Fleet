package main

// Lv2 - Portable plugin standard for audio systems
// Homepage: https://lv2plug.in/

import (
	"fmt"
	
	"os/exec"
)

func installLv2() {
	// Método 1: Descargar y extraer .tar.gz
	lv2_tar_url := "https://lv2plug.in/spec/lv2-1.18.10.tar.xz"
	lv2_cmd_tar := exec.Command("curl", "-L", lv2_tar_url, "-o", "package.tar.gz")
	err := lv2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lv2_zip_url := "https://lv2plug.in/spec/lv2-1.18.10.tar.xz"
	lv2_cmd_zip := exec.Command("curl", "-L", lv2_zip_url, "-o", "package.zip")
	err = lv2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lv2_bin_url := "https://lv2plug.in/spec/lv2-1.18.10.tar.xz"
	lv2_cmd_bin := exec.Command("curl", "-L", lv2_bin_url, "-o", "binary.bin")
	err = lv2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lv2_src_url := "https://lv2plug.in/spec/lv2-1.18.10.tar.xz"
	lv2_cmd_src := exec.Command("curl", "-L", lv2_src_url, "-o", "source.tar.gz")
	err = lv2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lv2_cmd_direct := exec.Command("./binary")
	err = lv2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
