package main

// Slirp4netns - User-mode networking for unprivileged network namespaces
// Homepage: https://github.com/rootless-containers/slirp4netns

import (
	"fmt"
	
	"os/exec"
)

func installSlirp4netns() {
	// Método 1: Descargar y extraer .tar.gz
	slirp4netns_tar_url := "https://github.com/rootless-containers/slirp4netns/archive/refs/tags/v1.3.1.tar.gz"
	slirp4netns_cmd_tar := exec.Command("curl", "-L", slirp4netns_tar_url, "-o", "package.tar.gz")
	err := slirp4netns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slirp4netns_zip_url := "https://github.com/rootless-containers/slirp4netns/archive/refs/tags/v1.3.1.zip"
	slirp4netns_cmd_zip := exec.Command("curl", "-L", slirp4netns_zip_url, "-o", "package.zip")
	err = slirp4netns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slirp4netns_bin_url := "https://github.com/rootless-containers/slirp4netns/archive/refs/tags/v1.3.1.bin"
	slirp4netns_cmd_bin := exec.Command("curl", "-L", slirp4netns_bin_url, "-o", "binary.bin")
	err = slirp4netns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slirp4netns_src_url := "https://github.com/rootless-containers/slirp4netns/archive/refs/tags/v1.3.1.src.tar.gz"
	slirp4netns_cmd_src := exec.Command("curl", "-L", slirp4netns_src_url, "-o", "source.tar.gz")
	err = slirp4netns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slirp4netns_cmd_direct := exec.Command("./binary")
	err = slirp4netns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
	fmt.Println("Instalando dependencia: jq")
	exec.Command("latte", "install", "jq").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libcap")
	exec.Command("latte", "install", "libcap").Run()
	fmt.Println("Instalando dependencia: libseccomp")
	exec.Command("latte", "install", "libseccomp").Run()
	fmt.Println("Instalando dependencia: libslirp")
	exec.Command("latte", "install", "libslirp").Run()
}
