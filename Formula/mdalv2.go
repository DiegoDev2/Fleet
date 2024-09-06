package main

// MdaLv2 - LV2 port of the MDA plugins
// Homepage: https://drobilla.net/software/mda-lv2.html

import (
	"fmt"
	
	"os/exec"
)

func installMdaLv2() {
	// Método 1: Descargar y extraer .tar.gz
	mdalv2_tar_url := "https://download.drobilla.net/mda-lv2-1.2.10.tar.xz"
	mdalv2_cmd_tar := exec.Command("curl", "-L", mdalv2_tar_url, "-o", "package.tar.gz")
	err := mdalv2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdalv2_zip_url := "https://download.drobilla.net/mda-lv2-1.2.10.tar.xz"
	mdalv2_cmd_zip := exec.Command("curl", "-L", mdalv2_zip_url, "-o", "package.zip")
	err = mdalv2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdalv2_bin_url := "https://download.drobilla.net/mda-lv2-1.2.10.tar.xz"
	mdalv2_cmd_bin := exec.Command("curl", "-L", mdalv2_bin_url, "-o", "binary.bin")
	err = mdalv2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdalv2_src_url := "https://download.drobilla.net/mda-lv2-1.2.10.tar.xz"
	mdalv2_cmd_src := exec.Command("curl", "-L", mdalv2_src_url, "-o", "source.tar.gz")
	err = mdalv2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdalv2_cmd_direct := exec.Command("./binary")
	err = mdalv2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sord")
	exec.Command("latte", "install", "sord").Run()
	fmt.Println("Instalando dependencia: lv2")
	exec.Command("latte", "install", "lv2").Run()
}
