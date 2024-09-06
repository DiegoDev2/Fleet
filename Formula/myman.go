package main

// Myman - Text-mode videogame inspired by Namco's Pac-Man
// Homepage: https://myman.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installMyman() {
	// Método 1: Descargar y extraer .tar.gz
	myman_tar_url := "https://downloads.sourceforge.net/project/myman/myman-cvs/myman-cvs-2009-10-30/myman-wip-2009-10-30.tar.gz"
	myman_cmd_tar := exec.Command("curl", "-L", myman_tar_url, "-o", "package.tar.gz")
	err := myman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	myman_zip_url := "https://downloads.sourceforge.net/project/myman/myman-cvs/myman-cvs-2009-10-30/myman-wip-2009-10-30.zip"
	myman_cmd_zip := exec.Command("curl", "-L", myman_zip_url, "-o", "package.zip")
	err = myman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	myman_bin_url := "https://downloads.sourceforge.net/project/myman/myman-cvs/myman-cvs-2009-10-30/myman-wip-2009-10-30.bin"
	myman_cmd_bin := exec.Command("curl", "-L", myman_bin_url, "-o", "binary.bin")
	err = myman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	myman_src_url := "https://downloads.sourceforge.net/project/myman/myman-cvs/myman-cvs-2009-10-30/myman-wip-2009-10-30.src.tar.gz"
	myman_cmd_src := exec.Command("curl", "-L", myman_src_url, "-o", "source.tar.gz")
	err = myman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	myman_cmd_direct := exec.Command("./binary")
	err = myman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: groff")
	exec.Command("latte", "install", "groff").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
