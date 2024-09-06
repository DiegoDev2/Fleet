package main

// Scs - Conic optimization via operator splitting
// Homepage: https://web.stanford.edu/~boyd/papers/scs.html

import (
	"fmt"
	
	"os/exec"
)

func installScs() {
	// Método 1: Descargar y extraer .tar.gz
	scs_tar_url := "https://github.com/cvxgrp/scs/archive/refs/tags/3.2.7.tar.gz"
	scs_cmd_tar := exec.Command("curl", "-L", scs_tar_url, "-o", "package.tar.gz")
	err := scs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scs_zip_url := "https://github.com/cvxgrp/scs/archive/refs/tags/3.2.7.zip"
	scs_cmd_zip := exec.Command("curl", "-L", scs_zip_url, "-o", "package.zip")
	err = scs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scs_bin_url := "https://github.com/cvxgrp/scs/archive/refs/tags/3.2.7.bin"
	scs_cmd_bin := exec.Command("curl", "-L", scs_bin_url, "-o", "binary.bin")
	err = scs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scs_src_url := "https://github.com/cvxgrp/scs/archive/refs/tags/3.2.7.src.tar.gz"
	scs_cmd_src := exec.Command("curl", "-L", scs_src_url, "-o", "source.tar.gz")
	err = scs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scs_cmd_direct := exec.Command("./binary")
	err = scs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
}
