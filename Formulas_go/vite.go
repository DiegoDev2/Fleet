package main

// Vite - Next generation frontend tooling. It's fast!
// Homepage: https://vitejs.dev/

import (
	"fmt"
	
	"os/exec"
)

func installVite() {
	// Método 1: Descargar y extraer .tar.gz
	vite_tar_url := "https://registry.npmjs.org/vite/-/vite-5.4.3.tgz"
	vite_cmd_tar := exec.Command("curl", "-L", vite_tar_url, "-o", "package.tar.gz")
	err := vite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vite_zip_url := "https://registry.npmjs.org/vite/-/vite-5.4.3.tgz"
	vite_cmd_zip := exec.Command("curl", "-L", vite_zip_url, "-o", "package.zip")
	err = vite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vite_bin_url := "https://registry.npmjs.org/vite/-/vite-5.4.3.tgz"
	vite_cmd_bin := exec.Command("curl", "-L", vite_bin_url, "-o", "binary.bin")
	err = vite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vite_src_url := "https://registry.npmjs.org/vite/-/vite-5.4.3.tgz"
	vite_cmd_src := exec.Command("curl", "-L", vite_src_url, "-o", "source.tar.gz")
	err = vite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vite_cmd_direct := exec.Command("./binary")
	err = vite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
