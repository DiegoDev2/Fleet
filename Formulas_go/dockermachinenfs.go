package main

// DockerMachineNfs - Activates NFS on docker-machine
// Homepage: https://github.com/adlogix/docker-machine-nfs

import (
	"fmt"
	
	"os/exec"
)

func installDockerMachineNfs() {
	// Método 1: Descargar y extraer .tar.gz
	dockermachinenfs_tar_url := "https://github.com/adlogix/docker-machine-nfs/archive/refs/tags/0.5.4.tar.gz"
	dockermachinenfs_cmd_tar := exec.Command("curl", "-L", dockermachinenfs_tar_url, "-o", "package.tar.gz")
	err := dockermachinenfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockermachinenfs_zip_url := "https://github.com/adlogix/docker-machine-nfs/archive/refs/tags/0.5.4.zip"
	dockermachinenfs_cmd_zip := exec.Command("curl", "-L", dockermachinenfs_zip_url, "-o", "package.zip")
	err = dockermachinenfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockermachinenfs_bin_url := "https://github.com/adlogix/docker-machine-nfs/archive/refs/tags/0.5.4.bin"
	dockermachinenfs_cmd_bin := exec.Command("curl", "-L", dockermachinenfs_bin_url, "-o", "binary.bin")
	err = dockermachinenfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockermachinenfs_src_url := "https://github.com/adlogix/docker-machine-nfs/archive/refs/tags/0.5.4.src.tar.gz"
	dockermachinenfs_cmd_src := exec.Command("curl", "-L", dockermachinenfs_src_url, "-o", "source.tar.gz")
	err = dockermachinenfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockermachinenfs_cmd_direct := exec.Command("./binary")
	err = dockermachinenfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
