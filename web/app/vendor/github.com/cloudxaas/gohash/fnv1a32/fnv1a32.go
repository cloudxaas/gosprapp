package fnv1a32

const (
    FNV_offset_basis uint32 = 2166136261
    FNV_prime        uint32 = 16777619
)

func Hash(data []byte) uint32 {    
    var hash uint32 = FNV_offset_basis
    for _, b := range data {
        hash ^= uint32(b)
        hash *= FNV_prime
    }
    return hash
}
