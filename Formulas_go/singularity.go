package main

// Singularity - Application container and unprivileged sandbox platform for Linux
// Homepage: https://apptainer.org/

import (
	"fmt"
	
	"os/exec"
)

func installSingularity() {
	// Método 1: Descargar y extraer .tar.gz
	singularity_tar_url := "https://github.com/apptainer/singularity/releases/download/v3.8.7/singularity-3.8.7.tar.gz"
	singularity_cmd_tar := exec.Command("curl", "-L", singularity_tar_url, "-o", "package.tar.gz")
	err := singularity_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	singularity_zip_url := "https://github.com/apptainer/singularity/releases/download/v3.8.7/singularity-3.8.7.zip"
	singularity_cmd_zip := exec.Command("curl", "-L", singularity_zip_url, "-o", "package.zip")
	err = singularity_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	singularity_bin_url := "https://github.com/apptainer/singularity/releases/download/v3.8.7/singularity-3.8.7.bin"
	singularity_cmd_bin := exec.Command("curl", "-L", singularity_bin_url, "-o", "binary.bin")
	err = singularity_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	singularity_src_url := "https://github.com/apptainer/singularity/releases/download/v3.8.7/singularity-3.8.7.src.tar.gz"
	singularity_cmd_src := exec.Command("curl", "-L", singularity_src_url, "-o", "source.tar.gz")
	err = singularity_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	singularity_cmd_direct := exec.Command("./binary")
	err = singularity_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libseccomp")
exec.Command("latte", "install", "libseccomp")
	fmt.Println("Instalando dependencia: squashfs")
exec.Command("latte", "install", "squashfs")
}
