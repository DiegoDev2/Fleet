package main

// Nftables - Netfilter tables userspace tools
// Homepage: https://netfilter.org/projects/nftables/

import (
	"fmt"
	
	"os/exec"
)

func installNftables() {
	// Método 1: Descargar y extraer .tar.gz
	nftables_tar_url := "https://www.netfilter.org/pub/nftables/nftables-1.1.0.tar.xz"
	nftables_cmd_tar := exec.Command("curl", "-L", nftables_tar_url, "-o", "package.tar.gz")
	err := nftables_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nftables_zip_url := "https://www.netfilter.org/pub/nftables/nftables-1.1.0.tar.xz"
	nftables_cmd_zip := exec.Command("curl", "-L", nftables_zip_url, "-o", "package.zip")
	err = nftables_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nftables_bin_url := "https://www.netfilter.org/pub/nftables/nftables-1.1.0.tar.xz"
	nftables_cmd_bin := exec.Command("curl", "-L", nftables_bin_url, "-o", "binary.bin")
	err = nftables_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nftables_src_url := "https://www.netfilter.org/pub/nftables/nftables-1.1.0.tar.xz"
	nftables_cmd_src := exec.Command("curl", "-L", nftables_src_url, "-o", "source.tar.gz")
	err = nftables_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nftables_cmd_direct := exec.Command("./binary")
	err = nftables_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: libedit")
	exec.Command("latte", "install", "libedit").Run()
	fmt.Println("Instalando dependencia: libmnl")
	exec.Command("latte", "install", "libmnl").Run()
	fmt.Println("Instalando dependencia: libnftnl")
	exec.Command("latte", "install", "libnftnl").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
