package main

// BtrfsProgs - Userspace utilities to manage btrfs filesystems
// Homepage: https://btrfs.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installBtrfsProgs() {
	// Método 1: Descargar y extraer .tar.gz
	btrfsprogs_tar_url := "https://mirrors.edge.kernel.org/pub/linux/kernel/people/kdave/btrfs-progs/btrfs-progs-v6.10.1.tar.xz"
	btrfsprogs_cmd_tar := exec.Command("curl", "-L", btrfsprogs_tar_url, "-o", "package.tar.gz")
	err := btrfsprogs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	btrfsprogs_zip_url := "https://mirrors.edge.kernel.org/pub/linux/kernel/people/kdave/btrfs-progs/btrfs-progs-v6.10.1.tar.xz"
	btrfsprogs_cmd_zip := exec.Command("curl", "-L", btrfsprogs_zip_url, "-o", "package.zip")
	err = btrfsprogs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	btrfsprogs_bin_url := "https://mirrors.edge.kernel.org/pub/linux/kernel/people/kdave/btrfs-progs/btrfs-progs-v6.10.1.tar.xz"
	btrfsprogs_cmd_bin := exec.Command("curl", "-L", btrfsprogs_bin_url, "-o", "binary.bin")
	err = btrfsprogs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	btrfsprogs_src_url := "https://mirrors.edge.kernel.org/pub/linux/kernel/people/kdave/btrfs-progs/btrfs-progs-v6.10.1.tar.xz"
	btrfsprogs_cmd_src := exec.Command("curl", "-L", btrfsprogs_src_url, "-o", "source.tar.gz")
	err = btrfsprogs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	btrfsprogs_cmd_direct := exec.Command("./binary")
	err = btrfsprogs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: e2fsprogs")
exec.Command("latte", "install", "e2fsprogs")
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
