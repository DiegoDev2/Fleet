package main

// Cutter - Unit Testing Framework for C and C++
// Homepage: https://cutter.osdn.jp/

import (
	"fmt"
	
	"os/exec"
)

func installCutter() {
	// Método 1: Descargar y extraer .tar.gz
	cutter_tar_url := "https://osdn.mirror.constant.com/cutter/73761/cutter-1.2.8.tar.gz"
	cutter_cmd_tar := exec.Command("curl", "-L", cutter_tar_url, "-o", "package.tar.gz")
	err := cutter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cutter_zip_url := "https://osdn.mirror.constant.com/cutter/73761/cutter-1.2.8.zip"
	cutter_cmd_zip := exec.Command("curl", "-L", cutter_zip_url, "-o", "package.zip")
	err = cutter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cutter_bin_url := "https://osdn.mirror.constant.com/cutter/73761/cutter-1.2.8.bin"
	cutter_cmd_bin := exec.Command("curl", "-L", cutter_bin_url, "-o", "binary.bin")
	err = cutter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cutter_src_url := "https://osdn.mirror.constant.com/cutter/73761/cutter-1.2.8.src.tar.gz"
	cutter_cmd_src := exec.Command("curl", "-L", cutter_src_url, "-o", "source.tar.gz")
	err = cutter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cutter_cmd_direct := exec.Command("./binary")
	err = cutter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
