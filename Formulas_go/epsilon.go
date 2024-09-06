package main

// Epsilon - Powerful wavelet image compressor
// Homepage: https://sourceforge.net/projects/epsilon-project/

import (
	"fmt"
	
	"os/exec"
)

func installEpsilon() {
	// Método 1: Descargar y extraer .tar.gz
	epsilon_tar_url := "https://downloads.sourceforge.net/project/epsilon-project/epsilon/0.9.2/epsilon-0.9.2.tar.gz"
	epsilon_cmd_tar := exec.Command("curl", "-L", epsilon_tar_url, "-o", "package.tar.gz")
	err := epsilon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	epsilon_zip_url := "https://downloads.sourceforge.net/project/epsilon-project/epsilon/0.9.2/epsilon-0.9.2.zip"
	epsilon_cmd_zip := exec.Command("curl", "-L", epsilon_zip_url, "-o", "package.zip")
	err = epsilon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	epsilon_bin_url := "https://downloads.sourceforge.net/project/epsilon-project/epsilon/0.9.2/epsilon-0.9.2.bin"
	epsilon_cmd_bin := exec.Command("curl", "-L", epsilon_bin_url, "-o", "binary.bin")
	err = epsilon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	epsilon_src_url := "https://downloads.sourceforge.net/project/epsilon-project/epsilon/0.9.2/epsilon-0.9.2.src.tar.gz"
	epsilon_cmd_src := exec.Command("curl", "-L", epsilon_src_url, "-o", "source.tar.gz")
	err = epsilon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	epsilon_cmd_direct := exec.Command("./binary")
	err = epsilon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: popt")
exec.Command("latte", "install", "popt")
}
