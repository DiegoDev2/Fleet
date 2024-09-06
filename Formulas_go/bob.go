package main

// Bob - Version manager for neovim
// Homepage: https://github.com/MordechaiHadad/bob

import (
	"fmt"
	
	"os/exec"
)

func installBob() {
	// Método 1: Descargar y extraer .tar.gz
	bob_tar_url := "https://github.com/MordechaiHadad/bob/archive/refs/tags/v4.0.1.tar.gz"
	bob_cmd_tar := exec.Command("curl", "-L", bob_tar_url, "-o", "package.tar.gz")
	err := bob_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bob_zip_url := "https://github.com/MordechaiHadad/bob/archive/refs/tags/v4.0.1.zip"
	bob_cmd_zip := exec.Command("curl", "-L", bob_zip_url, "-o", "package.zip")
	err = bob_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bob_bin_url := "https://github.com/MordechaiHadad/bob/archive/refs/tags/v4.0.1.bin"
	bob_cmd_bin := exec.Command("curl", "-L", bob_bin_url, "-o", "binary.bin")
	err = bob_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bob_src_url := "https://github.com/MordechaiHadad/bob/archive/refs/tags/v4.0.1.src.tar.gz"
	bob_cmd_src := exec.Command("curl", "-L", bob_src_url, "-o", "source.tar.gz")
	err = bob_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bob_cmd_direct := exec.Command("./binary")
	err = bob_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
