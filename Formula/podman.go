package main

// Podman - Tool for managing OCI containers and pods
// Homepage: https://podman.io/

import (
	"fmt"
	
	"os/exec"
)

func installPodman() {
	// Método 1: Descargar y extraer .tar.gz
	podman_tar_url := "https://github.com/containers/podman.git"
	podman_cmd_tar := exec.Command("curl", "-L", podman_tar_url, "-o", "package.tar.gz")
	err := podman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	podman_zip_url := "https://github.com/containers/podman.git"
	podman_cmd_zip := exec.Command("curl", "-L", podman_zip_url, "-o", "package.zip")
	err = podman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	podman_bin_url := "https://github.com/containers/podman.git"
	podman_cmd_bin := exec.Command("curl", "-L", podman_bin_url, "-o", "binary.bin")
	err = podman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	podman_src_url := "https://github.com/containers/podman.git"
	podman_cmd_src := exec.Command("curl", "-L", podman_src_url, "-o", "source.tar.gz")
	err = podman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	podman_cmd_direct := exec.Command("./binary")
	err = podman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: go-md2man")
	exec.Command("latte", "install", "go-md2man").Run()
	fmt.Println("Instalando dependencia: make")
	exec.Command("latte", "install", "make").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: conmon")
	exec.Command("latte", "install", "conmon").Run()
	fmt.Println("Instalando dependencia: crun")
	exec.Command("latte", "install", "crun").Run()
	fmt.Println("Instalando dependencia: fuse-overlayfs")
	exec.Command("latte", "install", "fuse-overlayfs").Run()
	fmt.Println("Instalando dependencia: gpgme")
	exec.Command("latte", "install", "gpgme").Run()
	fmt.Println("Instalando dependencia: libseccomp")
	exec.Command("latte", "install", "libseccomp").Run()
	fmt.Println("Instalando dependencia: passt")
	exec.Command("latte", "install", "passt").Run()
	fmt.Println("Instalando dependencia: slirp4netns")
	exec.Command("latte", "install", "slirp4netns").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
