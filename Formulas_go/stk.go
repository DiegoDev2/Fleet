package main

// Stk - Sound Synthesis Toolkit
// Homepage: https://ccrma.stanford.edu/software/stk/

import (
	"fmt"
	
	"os/exec"
)

func installStk() {
	// Método 1: Descargar y extraer .tar.gz
	stk_tar_url := "https://ccrma.stanford.edu/software/stk/release/stk-5.0.1.tar.gz"
	stk_cmd_tar := exec.Command("curl", "-L", stk_tar_url, "-o", "package.tar.gz")
	err := stk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stk_zip_url := "https://ccrma.stanford.edu/software/stk/release/stk-5.0.1.zip"
	stk_cmd_zip := exec.Command("curl", "-L", stk_zip_url, "-o", "package.zip")
	err = stk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stk_bin_url := "https://ccrma.stanford.edu/software/stk/release/stk-5.0.1.bin"
	stk_cmd_bin := exec.Command("curl", "-L", stk_bin_url, "-o", "binary.bin")
	err = stk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stk_src_url := "https://ccrma.stanford.edu/software/stk/release/stk-5.0.1.src.tar.gz"
	stk_cmd_src := exec.Command("curl", "-L", stk_src_url, "-o", "source.tar.gz")
	err = stk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stk_cmd_direct := exec.Command("./binary")
	err = stk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
}
