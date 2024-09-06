package main

// Pod2man - Perl documentation generator
// Homepage: https://www.eyrie.org/~eagle/software/podlators/

import (
	"fmt"
	
	"os/exec"
)

func installPod2man() {
	// Método 1: Descargar y extraer .tar.gz
	pod2man_tar_url := "https://archives.eyrie.org/software/perl/podlators-v6.0.2.tar.xz"
	pod2man_cmd_tar := exec.Command("curl", "-L", pod2man_tar_url, "-o", "package.tar.gz")
	err := pod2man_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pod2man_zip_url := "https://archives.eyrie.org/software/perl/podlators-v6.0.2.tar.xz"
	pod2man_cmd_zip := exec.Command("curl", "-L", pod2man_zip_url, "-o", "package.zip")
	err = pod2man_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pod2man_bin_url := "https://archives.eyrie.org/software/perl/podlators-v6.0.2.tar.xz"
	pod2man_cmd_bin := exec.Command("curl", "-L", pod2man_bin_url, "-o", "binary.bin")
	err = pod2man_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pod2man_src_url := "https://archives.eyrie.org/software/perl/podlators-v6.0.2.tar.xz"
	pod2man_cmd_src := exec.Command("curl", "-L", pod2man_src_url, "-o", "source.tar.gz")
	err = pod2man_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pod2man_cmd_direct := exec.Command("./binary")
	err = pod2man_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
