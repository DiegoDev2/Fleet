package main

// Iptables - Linux kernel packet control tool
// Homepage: https://www.netfilter.org/projects/iptables/index.html

import (
	"fmt"
	
	"os/exec"
)

func installIptables() {
	// Método 1: Descargar y extraer .tar.gz
	iptables_tar_url := "https://www.netfilter.org/pub/iptables/iptables-1.8.10.tar.xz"
	iptables_cmd_tar := exec.Command("curl", "-L", iptables_tar_url, "-o", "package.tar.gz")
	err := iptables_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iptables_zip_url := "https://www.netfilter.org/pub/iptables/iptables-1.8.10.tar.xz"
	iptables_cmd_zip := exec.Command("curl", "-L", iptables_zip_url, "-o", "package.zip")
	err = iptables_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iptables_bin_url := "https://www.netfilter.org/pub/iptables/iptables-1.8.10.tar.xz"
	iptables_cmd_bin := exec.Command("curl", "-L", iptables_bin_url, "-o", "binary.bin")
	err = iptables_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iptables_src_url := "https://www.netfilter.org/pub/iptables/iptables-1.8.10.tar.xz"
	iptables_cmd_src := exec.Command("curl", "-L", iptables_src_url, "-o", "source.tar.gz")
	err = iptables_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iptables_cmd_direct := exec.Command("./binary")
	err = iptables_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: linux-headers@5.15")
exec.Command("latte", "install", "linux-headers@5.15")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libmnl")
exec.Command("latte", "install", "libmnl")
	fmt.Println("Instalando dependencia: libnetfilter_conntrack")
exec.Command("latte", "install", "libnetfilter_conntrack")
	fmt.Println("Instalando dependencia: libnfnetlink")
exec.Command("latte", "install", "libnfnetlink")
	fmt.Println("Instalando dependencia: libnftnl")
exec.Command("latte", "install", "libnftnl")
	fmt.Println("Instalando dependencia: libpcap")
exec.Command("latte", "install", "libpcap")
	fmt.Println("Instalando dependencia: nftables")
exec.Command("latte", "install", "nftables")
}
