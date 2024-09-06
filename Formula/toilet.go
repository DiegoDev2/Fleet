package main

// Toilet - Color-based alternative to figlet (uses libcaca)
// Homepage: http://caca.zoy.org/wiki/toilet

import (
	"fmt"
	
	"os/exec"
)

func installToilet() {
	// Método 1: Descargar y extraer .tar.gz
	toilet_tar_url := "http://caca.zoy.org/raw-attachment/wiki/toilet/toilet-0.3.tar.gz"
	toilet_cmd_tar := exec.Command("curl", "-L", toilet_tar_url, "-o", "package.tar.gz")
	err := toilet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	toilet_zip_url := "http://caca.zoy.org/raw-attachment/wiki/toilet/toilet-0.3.zip"
	toilet_cmd_zip := exec.Command("curl", "-L", toilet_zip_url, "-o", "package.zip")
	err = toilet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	toilet_bin_url := "http://caca.zoy.org/raw-attachment/wiki/toilet/toilet-0.3.bin"
	toilet_cmd_bin := exec.Command("curl", "-L", toilet_bin_url, "-o", "binary.bin")
	err = toilet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	toilet_src_url := "http://caca.zoy.org/raw-attachment/wiki/toilet/toilet-0.3.src.tar.gz"
	toilet_cmd_src := exec.Command("curl", "-L", toilet_src_url, "-o", "source.tar.gz")
	err = toilet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	toilet_cmd_direct := exec.Command("./binary")
	err = toilet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libcaca")
	exec.Command("latte", "install", "libcaca").Run()
}
