package main

// Genext2fs - Generates an ext2 filesystem as a normal (non-root) user
// Homepage: https://genext2fs.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGenext2fs() {
	// Método 1: Descargar y extraer .tar.gz
	genext2fs_tar_url := "https://github.com/bestouff/genext2fs/archive/refs/tags/v1.5.0.tar.gz"
	genext2fs_cmd_tar := exec.Command("curl", "-L", genext2fs_tar_url, "-o", "package.tar.gz")
	err := genext2fs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	genext2fs_zip_url := "https://github.com/bestouff/genext2fs/archive/refs/tags/v1.5.0.zip"
	genext2fs_cmd_zip := exec.Command("curl", "-L", genext2fs_zip_url, "-o", "package.zip")
	err = genext2fs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	genext2fs_bin_url := "https://github.com/bestouff/genext2fs/archive/refs/tags/v1.5.0.bin"
	genext2fs_cmd_bin := exec.Command("curl", "-L", genext2fs_bin_url, "-o", "binary.bin")
	err = genext2fs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	genext2fs_src_url := "https://github.com/bestouff/genext2fs/archive/refs/tags/v1.5.0.src.tar.gz"
	genext2fs_cmd_src := exec.Command("curl", "-L", genext2fs_src_url, "-o", "source.tar.gz")
	err = genext2fs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	genext2fs_cmd_direct := exec.Command("./binary")
	err = genext2fs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
