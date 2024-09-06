package main

// Nuvie - Ultima 6 engine
// Homepage: https://nuvie.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installNuvie() {
	// Método 1: Descargar y extraer .tar.gz
	nuvie_tar_url := "https://downloads.sourceforge.net/project/nuvie/Nuvie/0.5/nuvie-0.5.tgz"
	nuvie_cmd_tar := exec.Command("curl", "-L", nuvie_tar_url, "-o", "package.tar.gz")
	err := nuvie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nuvie_zip_url := "https://downloads.sourceforge.net/project/nuvie/Nuvie/0.5/nuvie-0.5.tgz"
	nuvie_cmd_zip := exec.Command("curl", "-L", nuvie_zip_url, "-o", "package.zip")
	err = nuvie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nuvie_bin_url := "https://downloads.sourceforge.net/project/nuvie/Nuvie/0.5/nuvie-0.5.tgz"
	nuvie_cmd_bin := exec.Command("curl", "-L", nuvie_bin_url, "-o", "binary.bin")
	err = nuvie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nuvie_src_url := "https://downloads.sourceforge.net/project/nuvie/Nuvie/0.5/nuvie-0.5.tgz"
	nuvie_cmd_src := exec.Command("curl", "-L", nuvie_src_url, "-o", "source.tar.gz")
	err = nuvie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nuvie_cmd_direct := exec.Command("./binary")
	err = nuvie_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
}
