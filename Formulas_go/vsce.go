package main

// Vsce - Tool for packaging, publishing and managing VS Code extensions
// Homepage: https://code.visualstudio.com/api/working-with-extensions/publishing-extension#vsce

import (
	"fmt"
	
	"os/exec"
)

func installVsce() {
	// Método 1: Descargar y extraer .tar.gz
	vsce_tar_url := "https://registry.npmjs.org/vsce/-/vsce-2.15.0.tgz"
	vsce_cmd_tar := exec.Command("curl", "-L", vsce_tar_url, "-o", "package.tar.gz")
	err := vsce_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vsce_zip_url := "https://registry.npmjs.org/vsce/-/vsce-2.15.0.tgz"
	vsce_cmd_zip := exec.Command("curl", "-L", vsce_zip_url, "-o", "package.zip")
	err = vsce_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vsce_bin_url := "https://registry.npmjs.org/vsce/-/vsce-2.15.0.tgz"
	vsce_cmd_bin := exec.Command("curl", "-L", vsce_bin_url, "-o", "binary.bin")
	err = vsce_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vsce_src_url := "https://registry.npmjs.org/vsce/-/vsce-2.15.0.tgz"
	vsce_cmd_src := exec.Command("curl", "-L", vsce_src_url, "-o", "source.tar.gz")
	err = vsce_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vsce_cmd_direct := exec.Command("./binary")
	err = vsce_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libsecret")
exec.Command("latte", "install", "libsecret")
}
