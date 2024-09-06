package main

// Sheldon - Fast, configurable, shell plugin manager
// Homepage: https://sheldon.cli.rs

import (
	"fmt"
	
	"os/exec"
)

func installSheldon() {
	// Método 1: Descargar y extraer .tar.gz
	sheldon_tar_url := "https://github.com/rossmacarthur/sheldon/archive/refs/tags/0.8.0.tar.gz"
	sheldon_cmd_tar := exec.Command("curl", "-L", sheldon_tar_url, "-o", "package.tar.gz")
	err := sheldon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sheldon_zip_url := "https://github.com/rossmacarthur/sheldon/archive/refs/tags/0.8.0.zip"
	sheldon_cmd_zip := exec.Command("curl", "-L", sheldon_zip_url, "-o", "package.zip")
	err = sheldon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sheldon_bin_url := "https://github.com/rossmacarthur/sheldon/archive/refs/tags/0.8.0.bin"
	sheldon_cmd_bin := exec.Command("curl", "-L", sheldon_bin_url, "-o", "binary.bin")
	err = sheldon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sheldon_src_url := "https://github.com/rossmacarthur/sheldon/archive/refs/tags/0.8.0.src.tar.gz"
	sheldon_cmd_src := exec.Command("curl", "-L", sheldon_src_url, "-o", "source.tar.gz")
	err = sheldon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sheldon_cmd_direct := exec.Command("./binary")
	err = sheldon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: libgit2")
	exec.Command("latte", "install", "libgit2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
