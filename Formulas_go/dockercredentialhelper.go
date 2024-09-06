package main

// DockerCredentialHelper - Platform keystore credential helper for Docker
// Homepage: https://github.com/docker/docker-credential-helpers

import (
	"fmt"
	
	"os/exec"
)

func installDockerCredentialHelper() {
	// Método 1: Descargar y extraer .tar.gz
	dockercredentialhelper_tar_url := "https://github.com/docker/docker-credential-helpers/archive/refs/tags/v0.8.2.tar.gz"
	dockercredentialhelper_cmd_tar := exec.Command("curl", "-L", dockercredentialhelper_tar_url, "-o", "package.tar.gz")
	err := dockercredentialhelper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dockercredentialhelper_zip_url := "https://github.com/docker/docker-credential-helpers/archive/refs/tags/v0.8.2.zip"
	dockercredentialhelper_cmd_zip := exec.Command("curl", "-L", dockercredentialhelper_zip_url, "-o", "package.zip")
	err = dockercredentialhelper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dockercredentialhelper_bin_url := "https://github.com/docker/docker-credential-helpers/archive/refs/tags/v0.8.2.bin"
	dockercredentialhelper_cmd_bin := exec.Command("curl", "-L", dockercredentialhelper_bin_url, "-o", "binary.bin")
	err = dockercredentialhelper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dockercredentialhelper_src_url := "https://github.com/docker/docker-credential-helpers/archive/refs/tags/v0.8.2.src.tar.gz"
	dockercredentialhelper_cmd_src := exec.Command("curl", "-L", dockercredentialhelper_src_url, "-o", "source.tar.gz")
	err = dockercredentialhelper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dockercredentialhelper_cmd_direct := exec.Command("./binary")
	err = dockercredentialhelper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libsecret")
exec.Command("latte", "install", "libsecret")
}
