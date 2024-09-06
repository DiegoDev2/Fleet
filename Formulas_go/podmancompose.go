package main

// PodmanCompose - Alternative to docker-compose using podman
// Homepage: https://github.com/containers/podman-compose

import (
	"fmt"
	
	"os/exec"
)

func installPodmanCompose() {
	// Método 1: Descargar y extraer .tar.gz
	podmancompose_tar_url := "https://files.pythonhosted.org/packages/bd/67/0f8cf5ef346a22ce73dfdd0e60cf81342329b71a7fc118128929f0c07b62/podman_compose-1.2.0.tar.gz"
	podmancompose_cmd_tar := exec.Command("curl", "-L", podmancompose_tar_url, "-o", "package.tar.gz")
	err := podmancompose_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	podmancompose_zip_url := "https://files.pythonhosted.org/packages/bd/67/0f8cf5ef346a22ce73dfdd0e60cf81342329b71a7fc118128929f0c07b62/podman_compose-1.2.0.zip"
	podmancompose_cmd_zip := exec.Command("curl", "-L", podmancompose_zip_url, "-o", "package.zip")
	err = podmancompose_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	podmancompose_bin_url := "https://files.pythonhosted.org/packages/bd/67/0f8cf5ef346a22ce73dfdd0e60cf81342329b71a7fc118128929f0c07b62/podman_compose-1.2.0.bin"
	podmancompose_cmd_bin := exec.Command("curl", "-L", podmancompose_bin_url, "-o", "binary.bin")
	err = podmancompose_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	podmancompose_src_url := "https://files.pythonhosted.org/packages/bd/67/0f8cf5ef346a22ce73dfdd0e60cf81342329b71a7fc118128929f0c07b62/podman_compose-1.2.0.src.tar.gz"
	podmancompose_cmd_src := exec.Command("curl", "-L", podmancompose_src_url, "-o", "source.tar.gz")
	err = podmancompose_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	podmancompose_cmd_direct := exec.Command("./binary")
	err = podmancompose_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: podman")
exec.Command("latte", "install", "podman")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
