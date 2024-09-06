package main

// E2fsprogs - Utilities for the ext2, ext3, and ext4 file systems
// Homepage: https://e2fsprogs.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installE2fsprogs() {
	// Método 1: Descargar y extraer .tar.gz
	e2fsprogs_tar_url := "https://downloads.sourceforge.net/project/e2fsprogs/e2fsprogs/v1.47.1/e2fsprogs-1.47.1.tar.gz"
	e2fsprogs_cmd_tar := exec.Command("curl", "-L", e2fsprogs_tar_url, "-o", "package.tar.gz")
	err := e2fsprogs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	e2fsprogs_zip_url := "https://downloads.sourceforge.net/project/e2fsprogs/e2fsprogs/v1.47.1/e2fsprogs-1.47.1.zip"
	e2fsprogs_cmd_zip := exec.Command("curl", "-L", e2fsprogs_zip_url, "-o", "package.zip")
	err = e2fsprogs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	e2fsprogs_bin_url := "https://downloads.sourceforge.net/project/e2fsprogs/e2fsprogs/v1.47.1/e2fsprogs-1.47.1.bin"
	e2fsprogs_cmd_bin := exec.Command("curl", "-L", e2fsprogs_bin_url, "-o", "binary.bin")
	err = e2fsprogs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	e2fsprogs_src_url := "https://downloads.sourceforge.net/project/e2fsprogs/e2fsprogs/v1.47.1/e2fsprogs-1.47.1.src.tar.gz"
	e2fsprogs_cmd_src := exec.Command("curl", "-L", e2fsprogs_src_url, "-o", "source.tar.gz")
	err = e2fsprogs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	e2fsprogs_cmd_direct := exec.Command("./binary")
	err = e2fsprogs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
