package main

// Dsh - Dancer's shell, or distributed shell
// Homepage: https://www.netfort.gr.jp/~dancer/software/dsh.html.en

import (
	"fmt"
	
	"os/exec"
)

func installDsh() {
	// Método 1: Descargar y extraer .tar.gz
	dsh_tar_url := "https://www.netfort.gr.jp/~dancer/software/downloads/dsh-0.25.10.tar.gz"
	dsh_cmd_tar := exec.Command("curl", "-L", dsh_tar_url, "-o", "package.tar.gz")
	err := dsh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dsh_zip_url := "https://www.netfort.gr.jp/~dancer/software/downloads/dsh-0.25.10.zip"
	dsh_cmd_zip := exec.Command("curl", "-L", dsh_zip_url, "-o", "package.zip")
	err = dsh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dsh_bin_url := "https://www.netfort.gr.jp/~dancer/software/downloads/dsh-0.25.10.bin"
	dsh_cmd_bin := exec.Command("curl", "-L", dsh_bin_url, "-o", "binary.bin")
	err = dsh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dsh_src_url := "https://www.netfort.gr.jp/~dancer/software/downloads/dsh-0.25.10.src.tar.gz"
	dsh_cmd_src := exec.Command("curl", "-L", dsh_src_url, "-o", "source.tar.gz")
	err = dsh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dsh_cmd_direct := exec.Command("./binary")
	err = dsh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libdshconfig")
	exec.Command("latte", "install", "libdshconfig").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
}
