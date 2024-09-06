package main

// Sha1dc - Tool to detect SHA-1 collisions in files, including SHAttered
// Homepage: https://github.com/cr-marcstevens/sha1collisiondetection

import (
	"fmt"
	
	"os/exec"
)

func installSha1dc() {
	// Método 1: Descargar y extraer .tar.gz
	sha1dc_tar_url := "https://github.com/cr-marcstevens/sha1collisiondetection/archive/refs/tags/stable-v1.0.3.tar.gz"
	sha1dc_cmd_tar := exec.Command("curl", "-L", sha1dc_tar_url, "-o", "package.tar.gz")
	err := sha1dc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sha1dc_zip_url := "https://github.com/cr-marcstevens/sha1collisiondetection/archive/refs/tags/stable-v1.0.3.zip"
	sha1dc_cmd_zip := exec.Command("curl", "-L", sha1dc_zip_url, "-o", "package.zip")
	err = sha1dc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sha1dc_bin_url := "https://github.com/cr-marcstevens/sha1collisiondetection/archive/refs/tags/stable-v1.0.3.bin"
	sha1dc_cmd_bin := exec.Command("curl", "-L", sha1dc_bin_url, "-o", "binary.bin")
	err = sha1dc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sha1dc_src_url := "https://github.com/cr-marcstevens/sha1collisiondetection/archive/refs/tags/stable-v1.0.3.src.tar.gz"
	sha1dc_cmd_src := exec.Command("curl", "-L", sha1dc_src_url, "-o", "source.tar.gz")
	err = sha1dc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sha1dc_cmd_direct := exec.Command("./binary")
	err = sha1dc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
