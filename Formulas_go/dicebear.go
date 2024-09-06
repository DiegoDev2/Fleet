package main

// Dicebear - CLI for DiceBear - An avatar library for designers and developers
// Homepage: https://github.com/dicebear/dicebear

import (
	"fmt"
	
	"os/exec"
)

func installDicebear() {
	// Método 1: Descargar y extraer .tar.gz
	dicebear_tar_url := "https://registry.npmjs.org/dicebear/-/dicebear-9.2.1.tgz"
	dicebear_cmd_tar := exec.Command("curl", "-L", dicebear_tar_url, "-o", "package.tar.gz")
	err := dicebear_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dicebear_zip_url := "https://registry.npmjs.org/dicebear/-/dicebear-9.2.1.tgz"
	dicebear_cmd_zip := exec.Command("curl", "-L", dicebear_zip_url, "-o", "package.zip")
	err = dicebear_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dicebear_bin_url := "https://registry.npmjs.org/dicebear/-/dicebear-9.2.1.tgz"
	dicebear_cmd_bin := exec.Command("curl", "-L", dicebear_bin_url, "-o", "binary.bin")
	err = dicebear_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dicebear_src_url := "https://registry.npmjs.org/dicebear/-/dicebear-9.2.1.tgz"
	dicebear_cmd_src := exec.Command("curl", "-L", dicebear_src_url, "-o", "source.tar.gz")
	err = dicebear_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dicebear_cmd_direct := exec.Command("./binary")
	err = dicebear_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: vips")
exec.Command("latte", "install", "vips")
}
