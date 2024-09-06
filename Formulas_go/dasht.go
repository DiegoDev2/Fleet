package main

// Dasht - Search API docs offline, in your terminal or browser
// Homepage: https://sunaku.github.io/dasht

import (
	"fmt"
	
	"os/exec"
)

func installDasht() {
	// Método 1: Descargar y extraer .tar.gz
	dasht_tar_url := "https://github.com/sunaku/dasht/archive/refs/tags/v2.4.0.tar.gz"
	dasht_cmd_tar := exec.Command("curl", "-L", dasht_tar_url, "-o", "package.tar.gz")
	err := dasht_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dasht_zip_url := "https://github.com/sunaku/dasht/archive/refs/tags/v2.4.0.zip"
	dasht_cmd_zip := exec.Command("curl", "-L", dasht_zip_url, "-o", "package.zip")
	err = dasht_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dasht_bin_url := "https://github.com/sunaku/dasht/archive/refs/tags/v2.4.0.bin"
	dasht_cmd_bin := exec.Command("curl", "-L", dasht_bin_url, "-o", "binary.bin")
	err = dasht_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dasht_src_url := "https://github.com/sunaku/dasht/archive/refs/tags/v2.4.0.src.tar.gz"
	dasht_cmd_src := exec.Command("curl", "-L", dasht_src_url, "-o", "source.tar.gz")
	err = dasht_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dasht_cmd_direct := exec.Command("./binary")
	err = dasht_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: socat")
exec.Command("latte", "install", "socat")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: w3m")
exec.Command("latte", "install", "w3m")
	fmt.Println("Instalando dependencia: wget")
exec.Command("latte", "install", "wget")
}
